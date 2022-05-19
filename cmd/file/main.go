package file

import (
	"log"
	"net/http"

	"github.com/alice890308/blog-server/modules/file/service"
	"github.com/spf13/cobra"
)

func NewFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [service]",
		Short: "start file's service",
	}

	svc := service.NewService()
	fs := http.FileServer(http.Dir("/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/upload", svc.Upload)
	http.HandleFunc("/status", svc.Status)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

	return cmd
}
