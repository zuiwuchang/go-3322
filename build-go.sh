#!/bin/bash
#Program:
#       golang 自動編譯 腳本
#History:
#       2018-03-28 king first release
#		2018-08-30 tag commit date
#Email:
#       zuiwuchang@gmail.com

# 定義的 各種 輔助 函數
MkDir(){
	mkdir -p "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
MkOrClear(){
	if test -d "$1";then
		declare path="$1"
		path+="/*"
		rm "$path" -rf
		if [ "$?" != 0 ] ;then
			exit 1
		fi
	else
		MkDir $1
	fi
}
NewFile(){
	echo "$2" > "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}
WriteFile(){
	echo "$2" >> "$1"
	if [ "$?" != 0 ] ;then
		exit 1
	fi
}


CreateGoVersion(){
	MkDir version
	filename="version/version.go"
	package="version"

	# 返回 git 信息 時間
	tag=`git describe`
	if [ "$tag" == '' ];then
		tag="[unknown tag]"
	fi

	commit=`git rev-parse HEAD`
	if [ "$commit" == '' ];then
		commit="[unknow commit]"
	fi
	
	date=`date +'%Y-%m-%d %H:%M:%S'`

	# 打印 信息
	echo ${tag} $commit
	echo $date


	# 自動 創建 go 代碼
	NewFile $filename	"package $package"
	WriteFile $filename	''
	WriteFile $filename	'// Tag git tag'
	WriteFile $filename	"const Tag = \`$tag\`"
	WriteFile $filename	'// Commit git commit'
	WriteFile $filename	"const Commit = \`$commit\`"
	WriteFile $filename	'// Date build datetime'
	WriteFile $filename	"const Date = \`$date\`"
}

# 自動 創建 version.go 代碼
CreateGoVersion

# build
go build -ldflags "-s -w"
