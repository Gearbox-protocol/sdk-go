local NULL_ADDR = '0x0000000000000000000000000000000000000000';
{
  // CRV1 token and pool have same address: LUSD3CRV-f, FRAX3CRV-f, uses generic wrapper for getting underlying address for add/remove liquidity
  //  generic wrapper 0xA79828DF1850E8a3A3064576f380D90aECDD3359
  //
  // CRV2 token and pool have different contract: gusd curve,  add/remove liquidity has specific wrapper over the curve for getting in the underlying token of 3crv
  //
  // CRV3 token and pool have different contract: crvPlain3andSUSD, can directly withdraw and deposit from this contract, no wrapper
  //
  // CRV4 token and pool have different address: no wrapper and no base pool, non meta pools
  //
  tokens: {
    ETH: '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee',  // it is used at some places though ETH is not a token.
    ///////////////////////////
    // yearn curve pools/tokens
    ///////////////////////////
    yvCURVE_FRAX: '0xB4AdA607B9d6b2c9Ee07A275e9616B84AC560139',
    //
    yvDAI: '0xdA816459F1AB5631232FE5e97a05BBBb94970c95',
    yvUSDC: '0xa354F35829Ae975e850e23e9615b11Da1B3dC4DE',
    yvWETH: '0xa258C4606Ca8206D8aA700cE2143D7db854D168c',
    yvWBTC: '0xA696a63cc78DfFa1a63E9E50587C197387FF6C7E',
    yvSTETH: '0x15a2B3CfaFd696e1C783FE99eed168b78a3A371e',
    ///////////////////////////
    // convex curve pools
    ///////////////////////////
    cvx_FRAX3CRV_PHTM: NULL_ADDR,
    cvx_3CRV_PHTM: NULL_ADDR,
    cvx_GUSD3CRV_PHTM: NULL_ADDR,
    cvx_crvPlain3andSUSD_PHTM: NULL_ADDR,
    cvx_stETH_PHTM: NULL_ADDR,
    ///////////////////////////
    // curve meta pools
    ///////////////////////////
    FRAX3CRV: '0xd632f22692FaC7611d2AA1C0D552930D43CAEd3B',
    LUSD3CRV: '0xEd279fDD11cA84bEef15AF5D39BB4d4bEE23F0cA',
    GUSD3CRV: '0xD2967f45c4f384DEEa880F807Be904762a3DeA07',
    crvPlain3andSUSD: '0xc25a3a3b969415c80451098fa907ec722572917f',  // this is using a different underlying 3CRV pool 0xA5407eAE9Ba41422680e2e00537571bcC53efBfD
    ///////////////////////////
    // curve pools
    ///////////////////////////
    '3CRV': '0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490',
    steCRV: '0x06325440D014e39736583c165C2963BA99fAf14E',
    ////////////////////
    // tokens swappable on AMM
    ////////////////////
    FEI: '0x956F47F50A910163D8BF957Cf5846D573E7f87CA',
    FRAX: '0x853d955acef822db058eb8505911ed77f175b99e',
    SUSD: '0x57Ab1ec28D129707052df4dF418D58a2D46d5f51',
    stETH: '0xae7ab96520de3a18e5e111b5eaab095312d7fe84',
    GUSD: '0x056fd409e1d7a124bd7017459dfea2f387b6d5cd',
    LUSD: '0x5f98805A4E8be255a32880FDeC7F6728C6568bA0',
    // yellow
    SNX: '0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f',
    LQTY: '0x6DEA81C8171D0bA574754EF6F8b412F2Ed88c54D',
    SPELL: '0x090185f2135308BaD17527004364eBcC2D37e5F6',
    CVX: '0x4e3fbd56cd56c3e72c1403e103b45db9da5b9d2b',
    LDO: '0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32',
    FXS: '0x3432B6A60D23Ca0dFCa7761B7ab56459D9C964D0',
    CRV: '0xD533a949740bb3306d119CC777fa900bA034cd52',
    // red
    SUSHI: '0x6b3595068778dd592e39a122f4f5a5cf09c90fe2',
    UNI: '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
    '1INCH': '0x111111111117dc0aa78b770fa6a738034120c302',
    YFI: '0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e',
    FTM: '0x4E15361FD6b4BB609Fa63C81A2be19d873717870',
    AAVE: '0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9',
    DPI: '0x1494ca1f11d487c2bbe4543e90080aeba4ba3c2b',
    LUNA: '0xd2877702675e6ceb975b4a1dff9fb7baf4c91ea9',
    COMP: '0xc00e94cb662c3520282e6f5717214004a7f26888',
    LINK: '0x514910771AF9Ca656af840dff83E8264EcF986CA',
    //////////////////////
    // underlying asset
    //////////////////////
    USDC: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    DAI: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
    WETH: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2',
    WBTC: '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
    //
    USDT: '0xdac17f958d2ee523a2206206994597c13d831ec7',
  },
  //
  //
  //||||||||||||||||||||||||||||||
  // exchanges
  //||||||||||||||||||||||||||||||
  //
  //
  ///////////////////////////
  // convex curve pools
  ///////////////////////////
  exchanges: {
    cvx_FRAX3CRV: '0xB900EF131301B307dB5eFcbed9DBb50A3e209B2e',
    cvx_3CRV: '0x689440f2Ff927E1f24c72F1087E1FAF471eCe1c8',
    cvx_stETH: '0x0A760466E1B4621579a82a39CB56Dda2F4E70f03',
    cvx_crvPlain3andSUSD: '0x22ee18aca7f3ee920d01f25da85840d12d98e8ca',
    cvx_GUSD3CRV: '0x7a7bbf95c44b144979360c3300b54a7d34b44985',
    ///////////////////////////
    // curve meta pools
    ///////////////////////////
    'LUSD3CRV-f': '0xEd279fDD11cA84bEef15AF5D39BB4d4bEE23F0cA',  // CRV1
    'FRAX3CRV-f': '0xd632f22692FaC7611d2AA1C0D552930D43CAEd3B',  //// CRV1
    GUSD3CRV_WRAPPER: '0x64448B78561690B70E17CBE8029a3e5c1bB7136e',  // CRV2 wrapper over curve pool 0x4f062658eaaf2c1ccf8c8e36d6824cdf41167956
    'crvPlain3andSUSD-f': '0xFCBa3E75865d2d561BE8D220616520c171F12851',  // CRV3
    ///////////////////////////
    // curve pools
    ///////////////////////////
    'steCRV-f': '0xDC24316b9AE028F1497c275EB9192a3Ea0f67022',  //CRV4
    '3CRV-f': '0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7',  //CRV4
    ///////////////////////////
    // AMM
    ///////////////////////////
    UNISWAPV2_ROUTER: '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D',
    UNISWAPV3_ROUTER: '0xE592427A0AEce92De3Edee1F18E0157C05861564',
    SUSHISWAP_ROUTER: '0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F',
  },
}
