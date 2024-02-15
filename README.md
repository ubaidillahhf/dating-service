# GoLang Dating Service

> **Author: Ubaidillah Hakim Fadly** \
> **Purpose: Case Study** \
> **Description:** \
> \
> *Our dating application service offers a streamlined experience, allowing users to register and login effortlessly. With intuitive swipe gestures, users can swiftly navigate through profiles, swiping left to pass or right to express interest. To ensure quality matches, users are presented with a curated selection of 10 profiles daily, eliminating repeat views. Additionally, users have the option to purchase a premium package, granting unlimited swipes and access to enhanced features for an enriched dating experience.*

### Running with docker composer

1. run
```bash
    docker composer up -d
```
2. app is running with port 8910, and ready to GO

### Running with live reload (for Development)

1. install [Air](https://github.com/cosmtrek/air)
2. create .air.toml at root (if not exist)
3. run
```bash
    air
```

### Running unit test (uncompleted yet bcs limited time, just user register for a sample)
```bash
    make test
```

### Prequisite tools
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Docker](https://www.docker.com/)
- [Golang](https://golang.org/)
- [DB Docs](https://dbdocs.io/docs)

### Documentation

- Access the DB documentation at [this address](https://dbdocs.io/ubed.dev/dating-service)
- Postman docs at [this address](https://documenter.getpostman.com/view/23782351/2sA2r6WPJE)

### Migration
1. Create New File
- Example:
```bash
    make new_migration name="create_users_table"
```
2. Migrate all
```bash
    make migrate_up
```
### Deployment
1. CI proses using jenkins, look at Jenkinsfile
2. CD proses using push method trigger from webhook dockerhub to endpoint generated by portainer
