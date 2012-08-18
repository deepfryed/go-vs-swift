# Go vs Swift

## Setup

```
go get github.com/bmizerany/pq
```

```
bundle install
```

## Results

```
$ ./gopg 
pg #insert 209.812ms
pg #select 826.244ms

$ ./swift.rb 
                           user     system      total        real
swift insert           0.020000   0.010000   0.030000 (  0.123857)
swift select           0.600000   0.020000   0.620000 (  0.753887)
```
