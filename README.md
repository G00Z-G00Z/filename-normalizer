# Filname-normilizer

A CLI app that handles normalizing the filenames in a folder. It can remove spaces, normalize all files to be snakecase, or cammel case.

# Development

To start developing, use the following commands to create a docker container with a golang enviroment.

```bash
# This command builds the docker image
docker build -t filename-normilizer --target dev .

# This commands runs an interactive container with go inside
docker run -it --rm -v ${PWD}:/app filename-normilizer bash
```
