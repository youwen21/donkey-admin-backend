package server

import (
	"io"
	"io/fs"
	"net/http"
)

//http.StripPrefix("/admin/", http.FileServer(http.FS(sub)))

type StaticServer struct {
	Fs fs.FS

	ServerPath HistoryRouters

	fileServer http.Handler
}

func (s StaticServer) CanServe(urlPath string) bool {
	if s.ServerPath != nil && s.ServerPath.IsContain(urlPath) {
		return true
	}

	return false
}

func (s StaticServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 非文件，而是vue路由， 输出index.html首页
	if s.ServerPath != nil && s.ServerPath.IsContain(r.URL.Path) {
		// 文件不存在 输出首页
		fi, _ := s.Fs.Open("index.html")
		defer fi.Close()
		content, _ := io.ReadAll(fi)
		w.Write(content)
		return
	}

	// css, js, img等文件
	// 需要输出正确的header头，所以借助http.FileServer
	s.fileServer.ServeHTTP(w, r)
}

func NewStaticServer(distFs fs.FS) StaticServer {
	routers, _ := initPathRouters(distFs)
	return StaticServer{
		Fs:         distFs,
		ServerPath: routers,
		fileServer: http.FileServer(http.FS(distFs)),
	}
}

func initPathRouters(distFs fs.FS) (HistoryRouters, error) {
	hsRouters := make(HistoryRouters, 0)

	// 遍历 distFs 所有文件, 将文件路径添加到 hsRouters 中
	fs.WalkDir(distFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		hsRouters = append(hsRouters, path)
		return nil
	})

	return hsRouters, nil
}
