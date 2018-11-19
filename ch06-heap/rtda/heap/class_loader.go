package heap

import (
	"fmt"
)

type ClassLoader struct {
	cp       *classpath.ClassPath
	calssMap map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader{
	return &ClassLoader{
		cp : cp,
		calssMap : make(map[srring]*Class)
	}
}

func(self *ClassLoader) LoadClass(name string) *Class{
	if class, ok := self.calssMap[name]; ok{
		return class //类已经加载
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class{
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry) 
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry){
	data, entry, err := self.cp.ReadClass(name)
	if err !=nil{
		panic("java.lang.ClassNotFoundException:" +name)
	}
	return data, entry
}
