package classpath
import "os"
import "strings"
import "fmt"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator){
		fmt.Printf("classNamexxxxx1:%v\n",  "newCompositeEntry")
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*"){
		fmt.Printf("classNamexxxxx2:%v\n",  "newWildcardEntry")
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP"){
		fmt.Printf("classNamexxxxx3:%v\n",  "newZipEntry")
		return newZipEntry(path)
	}
	fmt.Printf("classNamexxxxx4:%v\n",  "newDirEntry")
	return newDirEntry(path)
}
