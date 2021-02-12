git pull
yarn build
pm2 delete awsl
pm2 start --name awsl yarn -- start

