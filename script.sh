find . \
  -type f \
  \( -name "*.go" -o -name "*.sql" -o -name "*.yml" -o -name "*.yaml" -o -name "Makefile" \) \
  ! -path "./.git/*" \
  ! -path "./vendor/*" \
  -print0 |
while IFS= read -r -d '' file; do
    echo "===== $file =====" >> all_code.txt
    cat "$file" >> all_code.txt
    echo -e "\n" >> all_code.txt
done