REPOS=(priceManager workflow priceSync workflowSync priceWeb)

for REPO in "${REPOS[@]}"
do
	echo "Starting connection ${REPO}\n"
	cd "../../${REPO}"
	if [ -e "app/main" ];
	then
  		go run app/main/main.go &
	fi
	cd ../infra/scripts/
continue
done

