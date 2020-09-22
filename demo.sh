#!/bin/bash

mv ../starwiz-customer/main ../starwiz-customer/main.old
mv customer ../starwiz-customer/main
if [$? -ne 0];then
  echo "deploy customer fail"
else
   docker-compose restart customer
fi


mv ../starwiz-system/main ../starwiz-system/main.old
mv system ../starwiz-system/main
if [$? -ne 0];then
   echo "deploy system fail"
else
  docker-compose restart system
fi

