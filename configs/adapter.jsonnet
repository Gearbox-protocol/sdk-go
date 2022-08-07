{
  YEARN_ADAPTER: [
    {
      stateMutability: 'view',
      type: 'function',
      name: 'pricePerShare',
      inputs: [],
      outputs: [
        {
          name: '',
          type: 'uint256',
        },
      ],
      gas: 18195,
    },
    {
      stateMutability: 'nonpayable',
      type: 'function',
      name: 'withdraw',
      inputs: [],
      outputs: [
        {
          name: '',
          type: 'uint256',
        },
      ],
    },
    {
      inputs: [],
      name: 'withdraw',
      outputs: [
        {
          internalType: 'uint256',
          name: '',
          type: 'uint256',
        },
      ],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      stateMutability: 'view',
      type: 'function',
      name: 'decimals',
      inputs: [],
      outputs: [
        {
          name: '',
          type: 'uint256',
        },
      ],
      gas: 3678,
    },
  ],
  UNISWAPV2_ADAPTER: [
    {
      inputs: [
        {
          internalType: 'uint256',
          name: 'amountIn',
          type: 'uint256',
        },
        {
          internalType: 'address[]',
          name: 'path',
          type: 'address[]',
        },
      ],
      name: 'getAmountsOut',
      outputs: [
        {
          internalType: 'uint256[]',
          name: 'amounts',
          type: 'uint256[]',
        },
      ],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        {
          internalType: 'uint256',
          name: 'rateMinRAY',
          type: 'uint256',
        },
        {
          internalType: 'address[]',
          name: 'path',
          type: 'address[]',
        },
        {
          internalType: 'uint256',
          name: 'deadline',
          type: 'uint256',
        },
      ],
      name: 'swapAllTokensForTokens',
      outputs: [
        {
          internalType: 'uint256[]',
          name: 'amounts',
          type: 'uint256[]',
        },
      ],
      stateMutability: 'nonpayable',
      type: 'function',
    },
  ],
  CURVE_ADAPTER: [
    {
      name: 'base_coins',
      outputs: [
        {
          type: 'address',
          name: '',
        },
      ],
      inputs: [
        {
          type: 'uint256',
          name: 'arg0',
        },
      ],
      stateMutability: 'view',
      type: 'function',
      gas: 1470,
    },
    {
      name: 'coins',
      outputs: [
        {
          type: 'address',
          name: '',
        },
      ],
      inputs: [
        {
          type: 'uint256',
          name: 'arg0',
        },
      ],
      stateMutability: 'view',
      type: 'function',
      gas: 1440,
    },
    {
      name: 'calc_withdraw_one_coin',
      outputs: [
        {
          type: 'uint256',
          name: '',
        },
      ],
      inputs: [
        {
          type: 'uint256',
          name: '_burn_amount',
        },
        {
          type: 'int128',
          name: 'i',
        },
      ],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        // {
        //   internalType: 'uint256',
        //   name: '',
        //   type: 'uint256',
        // },
        {
          internalType: 'int128',
          name: 'i',
          type: 'int128',
        },
        {
          internalType: 'uint256',
          name: '',
          type: 'uint256',
        },
      ],
      name: 'remove_all_liquidity_one_coin',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      name: 'get_virtual_price',
      outputs: [
        {
          type: 'uint256',
          name: '',
        },
      ],
      inputs: [],
      stateMutability: 'view',
      type: 'function',
      gas: 1011891,
    },
  ],
  CURVE_SUSD_ADAPTER: [
    {
      name: 'get_virtual_price',
      outputs: [
        {
          type: 'uint256',
          name: '',
        },
      ],
      inputs: [],
      stateMutability: 'view',
      type: 'function',
      gas: 1011891,
    },
    {
      name: 'coins',
      outputs: [
        {
          type: 'address',
          name: '',
        },
      ],
      inputs: [
        {
          type: 'int128',
          name: 'arg0',
        },
      ],
      constant: true,
      payable: false,
      type: 'function',
      gas: 1680,
    },
    {
      inputs: [
        // {
        //   internalType: 'uint256',
        //   name: '',
        //   type: 'uint256',
        // },
        {
          internalType: 'int128',
          name: 'i',
          type: 'int128',
        },
        {
          internalType: 'uint256',
          name: '',
          type: 'uint256',
        },
      ],
      name: 'remove_all_liquidity_one_coin',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      name: 'calc_withdraw_one_coin',
      outputs: [
        {
          type: 'uint256',
          name: '',
        },
      ],
      inputs: [
        {
          type: 'uint256',
          name: '_burn_amount',
        },
        {
          type: 'int128',
          name: 'i',
        },
      ],
      stateMutability: 'view',
      type: 'function',
    },
  ],
  CONVEX_ADAPTER: [],
  UNISWAPV3_ADAPTER: [
    {
      inputs: [
        {
          internalType: 'bytes',
          name: 'path',
          type: 'bytes',
        },
        {
          internalType: 'uint256',
          name: 'amountIn',
          type: 'uint256',
        },
      ],
      name: 'quoteExactInput',
      outputs: [
        {
          internalType: 'uint256',
          name: 'amountOut',
          type: 'uint256',
        },
      ],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        {
          components: [
            {
              internalType: 'bytes',
              name: 'path',
              type: 'bytes',
            },
            {
              internalType: 'uint256',
              name: 'deadline',
              type: 'uint256',
            },
            {
              internalType: 'uint256',
              name: 'rateMinRAY',
              type: 'uint256',
            },
          ],
          internalType: 'structIUniswapV3Adapter.ExactAllInputParams',
          name: 'params',
          type: 'tuple',
        },
      ],
      name: 'exactAllInput',
      outputs: [
        {
          internalType: 'uint256',
          name: 'amountOut',
          type: 'uint256',
        },
      ],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        {
          components: [
            {
              internalType: 'address',
              name: 'tokenIn',
              type: 'address',
            },
            {
              internalType: 'address',
              name: 'tokenOut',
              type: 'address',
            },
            {
              internalType: 'uint24',
              name: 'fee',
              type: 'uint24',
            },
            {
              internalType: 'uint256',
              name: 'deadline',
              type: 'uint256',
            },
            {
              internalType: 'uint256',
              name: 'rateMinRAY',
              type: 'uint256',
            },
            {
              internalType: 'uint160',
              name: 'sqrtPriceLimitX96',
              type: 'uint160',
            },
          ],
          internalType: 'structIUniswapV3Adapter.ExactAllInputSingleParams',
          name: 'params',
          type: 'tuple',
        },
      ],
      name: 'exactAllInputSingle',
      outputs: [
        {
          internalType: 'uint256',
          name: 'amountOut',
          type: 'uint256',
        },
      ],
      stateMutability: 'nonpayable',
      type: 'function',
    },
  ],
  CURVE_GENERIC_WRAPPER_ADAPTER: [],
}
