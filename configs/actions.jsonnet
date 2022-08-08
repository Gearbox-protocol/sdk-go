local NULL_ADDR = '0x0000000000000000000000000000000000000000';
local abi = import 'adapter.jsonnet';
local store = import 'mainnet.jsonnet';
local tokens = store.tokens;
local exchgs = store.exchanges;
local _red = [
  tokens.SUSHI,
  tokens.UNI,
  tokens['1INCH'],
  tokens.YFI,
  tokens.FTM,
  tokens.AAVE,
  tokens.DPI,
  tokens.LUNA,
  tokens.COMP,
  tokens.LINK,
];
local _yellow = [
  tokens.SNX,
  tokens.LQTY,
  tokens.SPELL,
  tokens.CVX,
  tokens.LDO,
  tokens.FXS,
  tokens.CRV,
];
local _green = [
  tokens.FEI,
  tokens.FRAX,
  tokens.SUSD,
  tokens.stETH,
  tokens.GUSD,
  // tokens.USDT,
  tokens.LUSD,
];
local _base = [
  tokens.USDC,
  tokens.WETH,
  tokens.DAI,
  tokens.WBTC,
];
local _3crv_tokens = [tokens.DAI, tokens.USDC, tokens.USDT];
local _non_synthetic_assets = _red + _yellow + _green + _base;
local swapDetails(inToken, exchanges, outTokens, transactContract=NULL_ADDR) = {
  inToken: inToken,
  exchanges: exchanges,
  outTokens: outTokens,
  transactContract: transactContract,
};

local arrayToObj(func, arr, init) =
  local aux(func, arr, running, idx) =
    if idx >= std.length(arr) then
      running
    else
      aux(func, arr, func(running, idx, arr[idx]), idx + 1) tailstrict;
  aux(func, arr, init, 0);

local mapFunc(running, idx, ele) =
  if std.isObject(ele) then
    running {
      [ele.inToken]: ele,
    }
  else if std.isString(ele) then
    running {
      [ele]: true,
    };

{
  adapters: {
    CONVEX_ADAPTER: {
      swapActions:: [
        swapDetails(tokens.cvx_FRAX3CRV_PHTM, [exchgs.cvx_FRAX3CRV], [tokens.FRAX3CRV]),
        swapDetails(tokens.cvx_3CRV_PHTM, exchgs.cvx_3CRV, [tokens['3CRV']]),
        swapDetails(tokens.cvx_stETH_PHTM, [exchgs.cvx_stETH], [tokens.stETH]),
        swapDetails(tokens.cvx_crvPlain3andSUSD_PHTM, [exchgs.cvx_crvPlain3andSUSD], [tokens.crvPlain3andSUSD]),
        swapDetails(tokens.cvx_GUSD3CRV_PHTM, [exchgs.cvx_GUSD3CRV], [tokens.GUSD3CRV]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.CONVEX_ADAPTER,
      name: 'ConvexAdapter',
    },
    YEARN_ADAPTER: {
      swapActions:: [
        swapDetails(tokens.yvCURVE_FRAX, [tokens.yvCURVE_FRAX], [tokens.FRAX3CRV]),
        swapDetails(tokens.yvDAI, [tokens.yvDAI], [tokens.DAI]),
        swapDetails(tokens.yvUSDC, [tokens.yvUSDC], [tokens.USDC]),
        swapDetails(tokens.yvWETH, [tokens.yvWETH], [tokens.WETH]),
        swapDetails(tokens.yvWBTC, [tokens.yvWBTC], [tokens.WBTC]),
        if tokens.yvSTETH != NULL_ADDR then swapDetails(tokens.yvSTETH, [tokens.yvSTETH], [tokens.stETH]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.YEARN_ADAPTER,
      name: 'YearnAdapter',
    },
    // CURVE_META_POOL_GENERIC_WRAPPER_ADAPTER: {
    //   swapActions:: [
    //     swapDetails(tokens.FRAX3CRV, [exchgs.FRAX3CRV_POOL], [tokens.FRAX] + _3crv_tokens),
    //     swapDetails(tokens.LUSD3CRV, [exchgs.LUSD3CRV_POOL], [tokens.LUSD] + _3crv_tokens),
    //   ],
    //   tokens: arrayToObj(mapFunc, self.swapActions, {}),
    //   abi: abi.CURVE_GENERIC_WRAPPER_ADAPTER,
    //   name: 'CurveGenericWrapperAdapter',
    // },
    CURVE_SUSD_ADAPTER: {
      // there are susd pool and  deposit contracts.
      // pool contract doesn't have add liquidity in one coin. So we are using deposit contract.
      swapActions:: [
        swapDetails(tokens.crvPlain3andSUSD, [exchgs.crvPlain3andSUSD_DEPOSIT], _3crv_tokens + [tokens.SUSD]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      // coins func on mainnet has arg of uint128 type
      abi: if store.network == 'kovan' then abi.CURVE_ADAPTER else abi.CURVE_SUSD_ADAPTER,
      name: 'CurveSUSDAdapter',
    },
    CURVE_3CRV_ADAPTER: {
      swapActions:: [
        swapDetails(tokens.GUSD3CRV, [exchgs.GUSD3CRV_POOL], [tokens.GUSD, tokens['3CRV']]),
        swapDetails(tokens.FRAX3CRV, [exchgs.FRAX3CRV_POOL], [tokens.FRAX, tokens['3CRV']]),
        swapDetails(tokens.LUSD3CRV, [exchgs.LUSD3CRV_POOL], [tokens.LUSD, tokens['3CRV']]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.CURVE_ADAPTER,
      name: 'Curve3crvAdapter',
    },
    CURVE_ADAPTER: {
      swapActions:: [
        ////////
        // gateway has the method for dealing with wrapper eth. whereas pool deals with native eth.
        // pool contract is needed for calc_withdraw_one_coin function.
        swapDetails(tokens.steCRV, [exchgs.steCRV_POOL], [tokens.WETH, tokens.stETH], exchgs.CURVE_STETH_GATEWAY),
        swapDetails(tokens['3CRV'], [exchgs['3CRV_POOL']], _3crv_tokens),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.CURVE_ADAPTER,
      name: 'CurveAdapter',
    },
    UNISWAPV2_ADAPTER: {
      tokens: arrayToObj(mapFunc, _non_synthetic_assets, {}),
      exchanges: [exchgs.SUSHISWAP_ROUTER, exchgs.UNISWAPV2_ROUTER],
      intermediaryTokens: [tokens.WETH, tokens.USDC, tokens.DAI, tokens.WBTC],
      abi: abi.UNISWAPV2_ADAPTER,
      name: 'UniswapV2Adapter',
    },
    UNISWAPV3_ADAPTER: {
      tokens: arrayToObj(mapFunc, _non_synthetic_assets, {}),
      exchanges: [exchgs.UNISWAPV3_ROUTER],
      Quoter: exchgs.UNISWAPV3_QUOTER,
      intermediaryTokens: [tokens.WETH, tokens.USDC, tokens.DAI, tokens.WBTC],
      abi: abi.UNISWAPV3_ADAPTER,
      name: 'UniswapV3Adapter',
    },
  },
}
