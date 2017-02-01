REPOS=(priceWeb priceSync priceManager priceManagerClient WorkflowClient workflowSync)

echo "${REPOS[@]}"
cd ../../
for REPO in "${REPOS[@]}"
do
	
      echo "cloning ${REPO}"
      git clone "https://github.com/RetailMarket/${REPO}.git"
continue
done

