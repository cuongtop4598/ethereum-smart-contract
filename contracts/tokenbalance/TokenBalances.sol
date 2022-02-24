// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0;

interface IERC20 {
  function balanceOf(address) external view returns (uint256);
}


contract DePocketBridge{
  function walletBalances(
    address userAddress,
    address[] calldata tokenAddresses
  ) external view returns (uint256[] memory) {
    require(tokenAddresses.length > 0);
    uint256[] memory balances = new uint256[](tokenAddresses.length);
    for (uint256 i = 0; i < tokenAddresses.length; i++) {
      if (tokenAddresses[i] != address(0x0)) {
        balances[i] = tokenBalance(userAddress, tokenAddresses[i]);
      } else {
        balances[i] = userAddress.balance;
      }
    }
    return balances;
  }
  /**
   * @notice Check the token balances of multiple wallets for multiple tokens
   * @dev return array and will fail if large arrays are inputted
   * @dev Returns array of token balances in base units
   * @param userAddresses address[]
   * @param tokenAddresses address[]
   * @return uint256[] token allowances array if possible
   */
  function allBalancesForManyAccounts(
    address[] calldata userAddresses,
    address[] calldata tokenAddresses
  ) external view returns (uint256[] memory) {
    uint256[] memory balances =
      new uint256[](tokenAddresses.length * userAddresses.length);
    for (uint256 user = 0; user < userAddresses.length; user++) {
      for (uint256 token = 0; token < tokenAddresses.length; token++) {
        if (tokenAddresses[token] != address(0x0)) {
          // ETH address in Etherdelta config
          balances[(user * tokenAddresses.length) + token] = tokenBalance(
            userAddresses[user],
            tokenAddresses[token]
          );
        } else {
          balances[(user * tokenAddresses.length) + token] = userAddresses[user]
            .balance;
        }
      }
    }
    return balances;
  }
  
  /**
   * @notice Check the token balance of a wallet in a token contract
   * @dev return 0 on returns 0 on invalid spender contract or non-contract address
   * @param userAddress address
   * @param tokenAddress address
   * @return uint256 token balance if possible
   */
  function tokenBalance(address userAddress, address tokenAddress)
    public
    view
    returns (uint256)
  {
    if (isContract(tokenAddress)) {
      IERC20 token = IERC20(tokenAddress);
      //  Check if balanceOf succeeds.
      (bool success, ) =
        address(token).staticcall(
          abi.encodeWithSelector(token.balanceOf.selector, userAddress)
        );
      if (success) {
        return token.balanceOf(userAddress);
      }
      return 0;
    }
    return 0;
  }
  
  // Not reliable, do not use when need strong verify
  function isContract(address addr) internal view returns (bool) {
    uint size;
    assembly { size := extcodesize(addr) }
    return size > 0;
  }
}