# expend-CIDR

`expend-CIDR` is a simple, fast command-line tool written in Go that expands a CIDR network range into a list of all usable IP addresses.

---

## Installation


```sh
go install -v github.com/h4xx00r/expend-CIDR@latest
```

## Usage

Simply provide a valid CIDR block as an argument. The tool will print each usable IP address to a new line.

```sh
./expend-CIDR <CIDR>
```

-----

## Example

Here's how to list all usable IPs in the `192.168.1.0/29` network.

**Command:**

```sh
./expend-CIDR 192.168.1.0/29
```

**Output:**

```
192.168.1.1
192.168.1.2
192.168.1.3
192.168.1.4
192.168.1.5
192.168.1.6
```
