# Filname-normilizer

A CLI app that handles normalizing the filenames in a folder. It can remove spaces, normalize all files to be snakecase, or cammel case.

## Available commands

- filename-normilizer normalize [type] [options]
    - type
        - snake_case (separates the words with underscores)
        - cammelCase (separates the words with uppercase letters)
    - options
        - r -> recursive
        - v -> verbose
        - i -> interactive


# Development

To start developing, use the following commands to create a docker container with a golang enviroment.

```bash
# This command builds the docker image
docker build -t filename-normilizer --target dev .

# This commands runs an interactive container with go inside
docker run -it --rm -v ${PWD}:/app filename-normilizer bash
```
