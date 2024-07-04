{
  constTokens: [
    {
      token: 'G-BLOCK',
      price: 0,
    },
    {
      token: 'G-OBS',
      price: 0,
    },
  ],
  arbBaseTokens: [
    'ARB',
    'GMX',
    'ezETH',
    'USDe',
  ],
  optBaseTokens: [
    'OP',
  ],
  baseTokens: [
    'FRAX',
    'SUSD',
    'crvUSD',
    'stETH',

    'GUSD',
    'LUSD',
    'SNX',
    'LQTY',
    'CVX',
    'LDO',
    'FXS',
    'CRV',
    '1INCH',
    'YFI',
    'AAVE',
    'DPI',
    'COMP',
    'LINK',
    'USDC',
    'DAI',
    'WETH',
    'WBTC',
    'USDT',
    'UNI',
    //
    'MKR',
    // 'BLUR',
    'RPL',
    'APE',
    'LDO',
    'GHO',
    //
    'PENDLE',  // on arbitrum
    'wstETH',
    'rETH',
    'cbETH',
    //
    'WLD',  // on optimism
  ],
  yearnTokens: [
    {
      token: 'yvDAI',
      underlying: 'DAI',
    },
    {
      token: 'yvUSDC',
      underlying: 'USDC',
    },
    {
      token: 'yvWETH',
      underlying: 'WETH',
    },
    {
      token: 'yvWBTC',
      underlying: 'WBTC',
    },
    {
      token: 'yvCurve_stETH',
      underlying: 'steCRV',
    },
    {
      token: 'sDAI',
      isMaker: true,
      underlying: 'DAI',
    },
  ],
  crvTokens: [
    {
      token: '3CRV',
      pool: '3CRV_POOL',
      underlyingTokens: [
        'DAI',
        'USDC',
        'USDT',
      ],
    },
    {
      token: 'FRAX3CRV',
      pool: 'FRAX3CRV_POOL',
      underlyingTokens: [
        'FRAX',
        '3CRV',
      ],
    },
    {
      token: 'steCRV',
      pool: 'steCRV_POOL',
      underlyingTokens: [
        'stETH',
        'WETH',
      ],
    },
    {
      token: 'crvPlain3andSUSD',
      pool: 'CURVE_SUSD_POOL',
      underlyingTokens: [
        'SUSD',
        'DAI',
        'USDC',
        'USDT',
      ],
    },
    {
      token: 'GUSD3CRV',
      pool: 'GUSD3CRV_POOL',
      underlyingTokens: [
        'GUSD',
        '3CRV',
      ],
    },
    {
      token: 'LUSD3CRV',
      pool: 'LUSD3CRV_POOL',
      underlyingTokens: [
        '3CRV',
        'LUSD',
      ],
    },
    {
      token: 'crvFRAX',
      pool: 'CURVE_FRAX_USDC_POOL',
      underlyingTokens: [
        'USDC',
        'FRAX',
      ],
    },
  ],
  copyPricesFor: {
    stkcvxFRAX3CRV: 'FRAX3CRV',
    cvxFRAX3CRV: 'FRAX3CRV',
    stkcvxsteCRV: 'steCRV',
    cvxsteCRV: 'steCRV',
    stkcvxgusd3CRV: 'GUSD3CRV',
    cvxgusd3CRV: 'GUSD3CRV',
    stkcvxcrvPlain3andSUSD: 'crvPlain3andSUSD',
    cvxcrvPlain3andSUSD: 'crvPlain3andSUSD',
    stkcvx3Crv: '3CRV',
    cvx3Crv: '3CRV',
    stkcvxLUSD3CRV: 'LUSD3CRV',
    cvxLUSD3CRV: 'LUSD3CRV',
    stkcvxcrvFRAX: 'crvFRAX',
  },
}
