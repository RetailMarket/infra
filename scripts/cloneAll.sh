REPOS=(priceWeb priceSync priceManager priceManagerClient WorkflowClient workflowSync workflow)

echo "${REPOS[@]}"
cd ../../
for REPO in "${REPOS[@]}"
do
	
      echo "cloning ${REPO}"
	if [ ! -d ${REPO} ];
	then
	      git clone "https://github.com/RetailMarket/${REPO}.git"
	fi
continue
done

