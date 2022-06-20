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
  ],
  CURVE_SUSD_ADAPTER: [
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
  ],
  CURVE_GENERIC_WRAPPER_ADAPTER: [],
}
