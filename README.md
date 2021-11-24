# kube-carbon-footprint

## UNDER CONSTRUCTION

## What is it?

KubeCarbonFootprint uses kubernetes metrics and calculates the estimated carbon footprint for a namespace

## How is it done ?

For the moment, the carbon footprint is calculated based on a formula:
```text
(CPU Usage usage) x (Cloud energy conversion factors [kWh]) x (Cloud provider Power Usage Effectiveness (PUE)) x (grid emissions factors [metric tons CO2e])
```
For the moment, we use static data for the PUE and cloud energy conversion factors.


