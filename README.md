# go-archiver


## Synopsis

Helps to archive files.

## Details
  
### Exported Functions

#### AddFileToArchive

  Description:
    Add source file to the end of tar archive. The sourcePath has to be the absolute path of the file to be added.
	Input:
		archivename string,
		sourcePath string,
		target string
	Returns:
	 error
	 
	 
#### CreateArchive
  Description:
    Create a tar file in the target path. If tar file already exists, it will return a pointer to te end of tar file easily append to
	Input: 
		archivename string,
		target string
	Returns:
		*tar.Writer, error 


## Contributors

Maria DeSouza <maria.g.desouza@gmail.com>

