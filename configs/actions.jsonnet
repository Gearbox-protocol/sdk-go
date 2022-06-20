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
local swapDetails(inToken, exchanges, outTokens) = {
  inToken: inToken,
  exchanges: exchanges,
  outTokens: outTokens,
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
    },
    YEARN_ADAPTER: {
      swapActions:: [
        swapDetails(tokens.yvCURVE_FRAX, [tokens.yvCURVE_FRAX], [tokens.FRAX3CRV]),
        swapDetails(tokens.yvDAI, [tokens.yvDAI], [tokens.DAI]),
        swapDetails(tokens.yvUSDC, [tokens.yvUSDC], [tokens.USDC]),
        swapDetails(tokens.yvWETH, [tokens.yvWETH], [tokens.WETH]),
        swapDetails(tokens.yvWBTC, [tokens.yvWBTC], [tokens.WBTC]),
        swapDetails(tokens.yvSTETH, [tokens.yvSTETH], [tokens.stETH]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.YEARN_ADAPTER,
    },
    CURVE_META_POOL_GENERIC_WRAPPER_ADAPTER: {
      swapActions:: [
        swapDetails(tokens.FRAX3CRV, [exchgs['FRAX3CRV-f']], [tokens.FRAX] + _3crv_tokens),
        swapDetails(tokens.LUSD3CRV, [exchgs['LUSD3CRV-f']], [tokens.LUSD] + _3crv_tokens),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.CURVE_GENERIC_WRAPPER_ADAPTER,
    },
    CURVE_SUSD_ADAPTER: {
      swapActions:: [
        swapDetails(tokens.crvPlain3andSUSD, [exchgs['crvPlain3andSUSD-f']], _3crv_tokens + [tokens.SUSD]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.CURVE_SUSD_ADAPTER,
    },
    CURVE_ADAPTER: {
      swapActions:: [
        // Metapools
        swapDetails(tokens.FRAX3CRV, [exchgs['FRAX3CRV-f']], [tokens.FRAX, tokens['3CRV']]),
        swapDetails(tokens.LUSD3CRV, [exchgs['LUSD3CRV-f']], [tokens.LUSD, tokens['3CRV']]),
        //
        swapDetails(tokens.GUSD3CRV, [exchgs.GUSD3CRV_WRAPPER], [tokens.GUSD] + _3crv_tokens),
        ////////
        swapDetails(tokens['3CRV'], [exchgs['3CRV-f']], _3crv_tokens),
        swapDetails(tokens.steCRV, [exchgs['steCRV-f']], [tokens.ETH, tokens.stETH]),
      ],
      tokens: arrayToObj(mapFunc, self.swapActions, {}),
      abi: abi.CURVE_ADAPTER,
    },
    UNISWAPV2_ADAPTER: {
      tokens: arrayToObj(mapFunc, _non_synthetic_assets, {}),
      exchanges: [exchgs.SUSHISWAP_ROUTER, exchgs.UNISWAPV2_ROUTER],
      intermediaryTokens: [tokens.WETH, tokens.USDC, tokens.DAI, tokens.WBTC],
      abi: abi.UNISWAPV2_ADAPTER,
    },
    UNISWAPV3_ADAPTER: {
      tokens: arrayToObj(mapFunc, _non_synthetic_assets, {}),
      exchanges: [exchgs.UNISWAPV3_ROUTER],
      intermediaryTokens: [tokens.WETH, tokens.USDC, tokens.DAI, tokens.WBTC],
      abi: abi.UNISWAPV3_ADAPTER,
    },
  },
}
