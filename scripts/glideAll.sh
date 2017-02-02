#REPOS="$(find ../.. -maxdepth 1 -mindepth 1 -type d)"

for REPO in ../../*
do
	echo "gliding ${REPO}"
	cd ${REPO}
	if [ -e "glide.yaml" ];
	then
  		glide cc; glide update; glide install;
	fi
	cd ../infra/scripts/
continue
done

