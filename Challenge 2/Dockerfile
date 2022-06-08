
FROM debian:11
RUN apt-get update && apt-get upgrade ;\ 
    useradd -ms /bin/bash your_first_name
USER your_first_name
WORKDIR /home/your_first_name
RUN touch test.txt
CMD ls -al