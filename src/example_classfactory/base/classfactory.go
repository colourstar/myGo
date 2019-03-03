package base

type Classer interface{
	Do()
}

type ClassFactory struct{
	classMap map[string] func()Classer
}

var classfactoryinst *ClassFactory

func (this *ClassFactory)RegClass(strRTTI string, funcinst func()Classer)  {
	this.classMap[strRTTI] = funcinst
}

func (this *ClassFactory)CreateClass(strRTTI string) Classer{
	funcinst := this.classMap[strRTTI]
	if funcinst != nil {
		return funcinst()
	}
	return nil
}

func GetInst() *ClassFactory  {
	if classfactoryinst == nil {
		classfactoryinst = &ClassFactory{}
		classfactoryinst.classMap = make(map[string] func() Classer)
	}

	return classfactoryinst
}