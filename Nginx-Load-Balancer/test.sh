for i in {1..5}
do
    curl --insecure -si https://localhost/ | grep -E "HTTP|Server"
    echo ""
done
