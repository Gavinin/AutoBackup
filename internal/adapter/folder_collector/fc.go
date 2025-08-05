package folder_collector

type IFolderCollector interface {
	GetFolderList([]string) ([]string, error)
}
