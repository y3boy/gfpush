.SILENT:
build:	gfpush.sh
	shc -f gfpush.sh
	mv gfpush.sh.x gfpush
	rm -f gfpush.sh.x.c
	echo "Build successfully!"
