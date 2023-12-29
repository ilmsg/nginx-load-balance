# nginx-load-balance
nginx ทำ load balance ด้วย round robin

run docker compose

    $ docker compose up -d --build

website: [http://webapp.local:3000/](http://webapp.local:3000/)

เรารัน webapp ไว้ 3 ตัว แต่ละตัวจะใส่ env ที่เป็น KEY เอาไว้เป็น key1, key2, key3 เพื่อให้รู้ว่าหน้าเว็บที่รันขึ้นมา ถูกเรียกจาก webapp ตัวไหน

เรากด refresh หน้าเว็บไปเรื่อยๆ จะเห็นว่าค่า key จะหมุนเปลี่ยนไปเรื่อยๆ key1 -> key2 -> key3 -> key1 .... หมุนวน แบบนี้ไปเรื่อยๆ

ถ้าเราลอง down webapp2 ลงไป หน้าเว็บก็จะออกแค่ key1 กับ key3 หมุนวนไป

![https://github.com/ilmsg/nginx-load-balance/blob/main/images/2023-12-29_130616.png?raw=true](https://github.com/ilmsg/nginx-load-balance/blob/main/images/2023-12-29_130616.png?raw=true)

![https://github.com/ilmsg/nginx-load-balance/blob/main/images/2023-12-29_130629.png?raw=true](https://github.com/ilmsg/nginx-load-balance/blob/main/images/2023-12-29_130629.png?raw=true)

![https://github.com/ilmsg/nginx-load-balance/blob/main/images/2023-12-29_130638.png?raw=true](https://github.com/ilmsg/nginx-load-balance/blob/main/images/2023-12-29_130638.png?raw=true)


note:
ตอนแรกจะทำ sticky round robin แต่พบว่า ต้องใช้ nginx ที่เป็น version nginx (>=1.5.7) ขึ้นไป ซึ่งนั้นจะเป็นแบบ commercial subscription