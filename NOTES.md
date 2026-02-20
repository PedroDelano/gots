# Chapter 1

## Elements of Exploratory Time Series Analysis

> The characteristic property of a time series is the fact that the data are not generated independently, their dispersion varies in time, they are often governed by a trend and they have cyclic components. Statistical procedures that suppose independent and identically distributed data are, therefore, excluded from the analysis of time series.

### The additive model

The additive model is the assumption that a time series is a realization of a random variable Yt that is the sum of four components:

- T[t]: trend
- Z[t]: non-random long term cyclic influence
- S[t]: non-random short term cyclic influenc
- R[t]: random variable

We can also use G[t] = T[t] + Z[t] to define the long term behavior of the time series.

### Models with a nonlinear trend

Consider Y[t] = T[t] + R[t] where the nonstochastic component is only the trend T[t]. Assuming E(R[t]) = 0, we have:

E(Y[t]) = T[t] := f(t)

A common assumption is that the function f dependes on several unkown parameters B1, B2, ..., Bn:

f(t) = f(t; B1, ..., Bn)

These parameters can be estimated from the set of realizations yt of the random variables Y[t]. A common approach is the least square estimate

```math
\sum_{t} \left(y_t - f(t;\hat{\beta}_1,\dots,\hat{\beta}_p)\right)^2 = \min_{\beta_1,\dots,\beta_p} \sum_{t} \left(y_t - f(t;\beta_1,\dots,\beta_p)\right)^2
```

The observed differences (yt - ŷt) are called residuals and can contain information about how well our model fits the data.
