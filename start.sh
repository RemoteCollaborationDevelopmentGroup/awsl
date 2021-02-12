git pull
npm run build
pm2 delete awsl
pm2 start --name awsl npm -- run start

