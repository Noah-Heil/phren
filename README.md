# Phren

Phren in Ancient Greek signifies "Mind"

## Purpose

Anyone with a mental illness is extrememly dependant on psychiatrist's. However pyschiatrist's are often overwhelmed, overworked and underinformed about their patients and what is going on in the health/medical/drug industry.

Phren bridges this gap between psychiatrists, their patients and the modern health industry. It is the memory that any normal human cannot have. It will allow normal people to get a personalized diagnosis of their pyscological disorders based uppon:

* Previous exeperiences with psyciatric medicine
* and if available Genomic data

When you have family memebers who are challeneged with serious mental illness and you watch them struggle everyday to function as a normal human being you begin to grasp at anything you can do to help. This is one of my attempts.

### Frustrations

In your journey to help your loved ones to find a normal, stable life you will probably run into some of the same things I have found like the following:

* I often see the helthcare industry be more concerened with making money off of other peoples pain and suffering...This is especially true when it comes to mental disorders.
* You will find that the resources available to you are minimal at best and nonexistant at worst.
* Often times your best option is to play a complex guessing game with a psychiatrist to try and figure out what medication fits just right for your loved one at this particular moment in time
* Heaven forbid that you want to start your family and want to have children
* Anything could change your biochemistry at anytime changing the effectiveness of your medication coctail that you have to take everyday just to function forcing you to start the whole complex guessing game to find a new mixture that works for you.
* You will often have no clue what the medications that your psychiatrist suggests actually do and you basically just pray it doesn't affect you adversely
* You will often wonder if the psychiatrist suggested the medication because they really think it is the best medicine for your case or because it is the first drug that came to their minds due to the excessively agressive marketing tactics used by drug manufactures.
* You will have no idea what side effects to look for from a new medication

## Getting Started

It doesn't do much right now but I am working towards it doing more. Currently it reads in a specific drug database into an internal data contruct.

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

You will need a functioning build of the Go language to hack away on Phren

[Install Golang](http://www.letmegooglethat.com/?q=how+to+install+the+go+language) -- How to Install the Go Language

You might also want to check out the Go Version Manager ([GVM](https://github.com/moovweb/gvm))

### Installing

A step by step series of examples that tell you how to get a development env running

Once you have the latest version of go installed you only have to clone this repository and your ready to start!

When you want to build Phren all you have to do is (from the root of the project) run:

```shell
go build
```

which will build the binary in your current directory.

I currently utilize the extremly useful [Cobra](https://github.com/spf13/cobra) package to build out the CLI for phren so to find a list of the flags and commands you can run you just have to run the newly build binary:

```shell
./phren
```

or just:

```shell
./phren run
```

## Running the tests

### Break down into end to end tests

Explain what these tests test and why

```shell
TBD
```

### And coding style tests

Explain what these tests test and why

```shell
TBD
```

## Deployment

Go applications are easy to run because they are usually compiled into simple quick binaries that you just have to build and drop onto a system to deploy.

## Built With

* golang

## Road Map

### Basic Functionality

- [x] Get Drug Data (Maybe a crawler? Or can I find it from somewhere?)
- [x] Load Drug Data (Found a Good Drug DB from www.drugbank.ca)
- [ ] Build Side Effect Data Type (Primative -- but take notes for advancing the Data type in the future)
- [ ] Map Out An Example Set of Side Effects
- [ ] Comparison Method
- [ ] Map Side Effects to Drug Targets

### Advanced Functionality

- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [x] completed

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* **Noah Heil** - *Initial Mess Maker and Bug Creator* - [Noah-Heil](https://github.com/Noah-Heil)

See also the list of [contributors](https://github.com/Noah-Heil/phren/contributors) who participated in this project.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* My Beautiful Wife and Family for all the sleepless nights and times I said I would make it to bed but didn't
* etc