# Requirements for build:
# gcc
# shc
.SILENT:
build:	gfpush.sh
	shc -f gfpush.sh
	mv gfpush.sh.x gfpush
	rm -f gfpush.sh.x.c
	sha256sum gfpush > sha256sum.txt
	echo "Build successfully!"
