
docker-compose build
docker-compose up -d

# Expected results are redirection of http://18.61.87.53:8080/ >> http://18.61.87.53:8081/srv2/




+--------+        +-------------+       +--------------------+
|        |        |             |       | Container Instance |
| Client +------->+  Service 1   +------>    1 (service1)     |
|        |        | (Kubernetes) |       +--------------------+
|        |        +-------------+
|        |                       
|        |                              
|        |                              
|        | ------------>> ENV NEEDED                             
|        |                                   
|        |
|        |        +-------------+       +--------------------+
|        |        |             |       | Container Instance |
|        +------->+  Service 2   +------>    1 (service2)    |
|        |        | (Kubernetes)|       +--------------------+
|        |        +-------------+  
|        |                         
|        |                                
|        |                                
|        |                                
|        |                                
+--------+
