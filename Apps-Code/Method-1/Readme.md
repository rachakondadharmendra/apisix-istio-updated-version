


docker build -t service1 -f Dockerfile.app1
docker run -d --name service1 -p 8080:8080 service1

# Expected results are redirection of http://18.61.87.53:8080/ >> http://18.61.87.53:8080/srv2/



+--------+        +-------------+       +--------------------+      +--------------------+
|        |        | Kubernetes  |       | Container Instance  |     | Container Instance |
| Client +------->+  Service     +------>    1 (api1)         |+--->|    2 (api2)        |
|        |        |  (service1) |       +--------------------+      +--------------------+
+--------+        +-------------+        
                  
