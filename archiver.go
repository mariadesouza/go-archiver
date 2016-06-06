package archiver

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// AddFileToArchive : Add source file to the end of tar archive. The sourcePath has to be the absolute path of the file to be added.
func AddFileToArchive(archivename, sourcePath, target string) ( error) {

	tw, err  := CreateArchive(archivename, target)
	if (err != nil)	{
		return err
	}

	file, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header, err := tar.FileInfoHeader(stat, "")
		if err != nil {
			return err
		}
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}


//CreateArchive : Will create/append archive in the target folder
/*	Input:
		archivename - name of tar
		target location of tar file
	Description: Create a tar file. If tar file already exists, it will return a pointer to te end of tar file easily append to 
*/
func CreateArchive(archivename, target string) (*tar.Writer, error) {

	tarfilename := filepath.Join(target, archivename)

	var tw *tar.Writer

	_, err := os.Stat(tarfilename)
	if os.IsNotExist(err) {
		tarfile, err := os.Create(tarfilename)
		if err != nil {	
			return nil, err
		}
		defer tarfile.Close()
		tw = tar.NewWriter(tarfile)
	} else { //Open the file to add more things to it

		f, err := os.OpenFile(tarfilename, os.O_RDWR, os.ModePerm)
		if err != nil {
			return nil, err
		}
		if _, err = f.Seek(-2<<9, os.SEEK_END); err != nil {
			return nil, err
		}
		tw = tar.NewWriter(f)
	}

	return tw, nil
}
