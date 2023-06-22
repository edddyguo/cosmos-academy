import type { ChainInfo } from "@keplr-wallet/types";
import React, { useEffect, useState } from "react";
import osmo from "../config/osmosis";
import { assertIsDeliverTxSuccess, SigningStargateClient } from "@cosmjs/stargate";

function Keplr() {
	const [chain, setChain] = useState<ChainInfo>(osmo);
	const [selected, setSelected] = useState<string>("OSMO");
	const [client, setClient] = useState<any>();
	const [address, setAddress] = useState<any>();

	const [balance, setBalance] = useState<any>();
	const [recipent, setRecipent] = useState<any>(
		"osmo1r9ufesd4ja09g4rcxxetpx675eu09m45q05wv7"
	);
	const [tx, setTx] = useState<any>();
	const [sendHash, setSendHash] = useState<any>();
	const [txRes, setTxRes] = useState<any>();

	// 初始化 chain
	useEffect(() => {
		connectWallet();
	}, [chain]);

	// 查余额
	useEffect(() => {
		if (!address && !client) return;
		getBalances();
	}, [address, client, sendHash]);

	// 连接keplr钱包  Todo
	const connectWallet = async () => {
		if (!window.keplr) {
			alert("Please install keplr wallet first");
		}

		await window.keplr.experimentalSuggestChain(chain);
		await window.keplr.enable(chain.chainId);

		let offlineSigner = window.keplr.getOfflineSigner(chain.chainId);
		let accounts = await offlineSigner.getAccounts();
		
		let client = await SigningStargateClient.connectWithSigner(
			chain.rpc,offlineSigner
		);

		setAddress(accounts[0].address);
		setClient(client);
	};

	// 余额查询  Done
	const getBalances = async () => {
		if (client) {
			let balance = client.getBalance(address, chain.stakeCurrency.coinDenom);
			setBalance(balance);
		}
	};

	// txhash查询  Done
	const getTx = async () => {
		if (tx && client) {
			let txRes = await client.getTx(address, sendHash);
			setTxRes(txRes);
		}
	};

	// 转账 Done
	const sendToken = async () => {
		if (!client || !recipent || !address) return;

		let decimal = 10 * 1e6;
		let amount = [
			{
				denom: chain.stakeCurrency.coinMinimalDenom,
				amount: decimal.toString()
			}
		];
		let fee = {
			amount: [
				{
					denom: chain.stakeCurrency.coinMinimalDenom,
					amount: 0.025
				}
			],
			gas: '200000'
		};

		try {
			const result = await client.sendTokens(
				address,
				recipent,
				amount,
				fee,
				""
			);
			assertIsDeliverTxSuccess(result);
			console.log(result);
			
			if (result.code == 0) {
				alert(
					"transfer success height:" + result.height +
					"hash:" + result.transactionHash
				);
				setTx(result.transactionHash);
			};
		} catch (e) {
			console.log(e);
		}
	};

	return (
		<div className="keplr">
			<h2>Keplr Wallet</h2>
			<label>
				<span>
					Chain: &nbsp;
					<select
						className="select"
						value={selected}
						onChange={(e) => setSelected(e.target.value)}
					>
						<option value="OSMO">OSMO</option>
						<option value="SPX">SPX</option>
					</select>
				</span>{" "}
				&nbsp;
				<button onClick={connectWallet}>
					{address ? "已连接" : "连接keplr"}
				</button>
			</label>
			<div className="weight">地址：{address}</div>
			<div className="weight">
				<span style={{ whiteSpace: "nowrap" }}>余额: &nbsp;</span>
				<div>
					{balance?.amount && (
						<>
							<span>
								{parseFloat(
									String(Number(balance?.amount) / Math.pow(10, 6))
								).toFixed(2)}
							</span>
							<span> {balance?.denom}</span>
						</>
					)}
				</div>
			</div>
			<hr />
			<label>1、sendTokens() & broadcastTx</label>
			<div>
				<input
					type="text"
					value={recipent}
					placeholder="address"
					style={{ width: "350px" }}
					onChange={(e) => setRecipent(e.target.value)}
				/>
				&nbsp;
				<button onClick={sendToken}>
					发送 10 {chain.feeCurrencies[0].coinMinimalDenom}
				</button>
			</div>
			<label>2、getTx()</label>
			<div>
				<input value={tx} readOnly style={{ width: "350px" }} />
				&nbsp;
				<button onClick={getTx}>查询</button>
			</div>
			<div className="tx">
				{txRes && (
					<>
						<div>height:{txRes?.height} </div>
						<div>gasUsed:{txRes?.gasUsed} </div>
						<div>gasWanted:{txRes?.gasWanted} </div>
					</>
				)}
			</div>
		</div>
	);
}

export default Keplr;
