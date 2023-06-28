# Cosmos SDK mint module
- There is no maximum supply.
- The inflation rate for each block is calculated based on the bonded ratio, the goal bonded percentage (67%), and the InflationRateChange parameter (which is set to a default value of 13%). The inflation rate is also bounded by the InflationMax and InflationMin values:

    ```
    // NextInflationRate returns the new inflation rate for the next hour.
    func (m Minter) NextInflationRate(params Params, bondedRatio sdk.Dec) sdk.Dec {
        // The target annual inflation rate is recalculated for each previsions cycle. The
        // inflation is also subject to a rate change (positive or negative) depending on
        // the distance from the desired ratio (67%). The maximum rate change possible is
        // defined to be 13% per year, however the annual inflation is capped as between
        // 7% and 20%.

        // (1 - bondedRatio/GoalBonded) * InflationRateChange
        inflationRateChangePerYear := sdk.OneDec().
            Sub(bondedRatio.Quo(params.GoalBonded)).
            Mul(params.InflationRateChange)
        inflationRateChange := inflationRateChangePerYear.Quo(sdk.NewDec(int64(params.BlocksPerYear)))

        // adjust the new annual inflation for this next cycle
        inflation := m.Inflation.Add(inflationRateChange) // note inflationRateChange may be negative
        if inflation.GT(params.InflationMax) {
            inflation = params.InflationMax
        }
        if inflation.LT(params.InflationMin) {
            inflation = params.InflationMin
        }

        return inflation
    }
    ```

- The number of newly minted coins each block is calculated by `currentSupply * inflationRate`:
    ```
    // NextAnnualProvisions returns the annual provisions based on current total
    // supply and inflation rate.
    func (m Minter) NextAnnualProvisions(_ Params, totalSupply sdk.Int) sdk.Dec {
	    remainingAnnualProvisions := uRioSupplyCap.Sub(totalSupply)
	    return m.Inflation.MulInt(remainingAnnualProvisions)
    }

    // BlockProvision returns the provisions for a block based on the annual
    // provisions rate.
    func (m Minter) BlockProvision(params Params) sdk.Coin {
	    provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))
	    return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
    }
    ```

- Simply explain:
    ```
    NewRate = OldRate + (1- bondedRatio/GoalBonded) *  params.InflationRateChange
	Reward =  NewRate * OldSupply
	NewSupply = OldSupply + Reward = (1 + NewRate) * OldSupply
    ```

    ```
    R(n+1) = R(n) + b
    ```
    - R is inflation rate at block n, InflationMin <= R <= InflationMax
    - b = (1 - bondedRatio / GoalBonded) * params.InflationRateChange | where GoaldBonded is 67% and default  and maximum of params.InflationRateChange is 13%.

    Therefore, we have:
    ```
    R(n+1) = R(n) + b1
    R(n+2) = R(n) + b1 + b2
    …
    R(n+k) = R(n) +b1 +b2 + … + bk

    => Reward(n+k) = Reward(n) * (1 + R(n) + b1) x … (1 + R(n) + b1 +  … + bk) x (R(n) +b1 + b2 + … + bk) / R(n)

    TotalSupply(n+k) = TotalSupply(n) x (1 + R(n) + b1) x … (1 + R(n) + b1 +  … + bk)
    ```
---

# Realio network mint module

- Maximum supply coins is 7.5e25 RIO:
    ```
    // Supply cap 1 RIO : 10^18 aRIO (attoRio)
    var rioSupplyCap, _ = math.NewIntFromString("75000000000000000000000000")
    ```
- Inflation rate is changed by SetParams function
- The number of newly minted coins each block: `(max supply - current supply) * inflation rate`
    ```
    // NextAnnualProvisions returns the annual provisions based on current unprovisioned total
    // supply and inflation rate.
    func (m Minter) NextAnnualProvisions(_ Params, totalSupply math.Int) sdk.Dec {
	    remainingAnnualProvisions := rioSupplyCap.Sub(totalSupply)
	    return m.Inflation.MulInt(remainingAnnualProvisions)
    }

    // BlockProvision returns the provisions for a block based on the annual
    // provisions rate.
    func (m Minter) BlockProvision(params Params) sdk.Coin {
	    provisionAmt := m.AnnualProvisions.QuoInt(math.NewInt(int64(params.BlocksPerYear)))
	    return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
    }
    ```
- Due to the non-utilization of the automatically calculated inflation rate, several functions and objects do not require unnecessary parameters such as InflationMax, InflationMin, GoalBonded, and so on:
    ```
    // realio-network
    func NewParams(
        mintDenom string, inflationRate sdk.Dec, blocksPerYear uint64,
    ) Params {
        return Params{
            MintDenom:     mintDenom,
            InflationRate: inflationRate,
            BlocksPerYear: blocksPerYear,
        }
    }

    // cosmos-sdk
    func NewParams(
        mintDenom string, inflationRateChange, inflationMax, inflationMin, goalBonded sdk.Dec, blocksPerYear uint64,
	) Params {

		return Params{
			MintDenom:           mintDenom,
			InflationRateChange: inflationRateChange,
			InflationMax:        inflationMax,
			InflationMin:        inflationMin,
			GoalBonded:          goalBonded,
			BlocksPerYear:       blocksPerYear,
		}
	}
    ```
