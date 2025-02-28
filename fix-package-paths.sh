PACKAGES=('common' 'enums' 'errors' 'resources' 'services')

function fix_package_path() {
    FILE=$1
    PACKAGE=$2
    MATCH="google.golang.org\/genproto\/googleapis\/ads\/googleads\/v11\/"
    REPLACE="github.com\/ercling\/google-ads-go\/"

    gsed -i "s/$MATCH$PACKAGE/$REPLACE$PACKAGE/g" $FILE
}

function fix_package_name() {
    FILE=$1
    PACKAGE=$2

    gsed -i "s/google_ads_googleads_v11_$PACKAGE/$PACKAGE/g" $FILE
}

echo "fixing packages"
for file in ./google/ads/googleads/v11/**/*.pb.go; do
    [ -e "$file" ] || continue
    for p in "${PACKAGES[@]}"; do
        fix_package_path $file $p
        fix_package_name $file $p
    done
done
mv ./google/ads/googleads/v11/resources/experiment_arm.pb.go ./google/ads/googleads/v11/resources/experiment_arm0.pb.go
mv ./google/ads/googleads/v11/* ./
echo "finished fixing packages"