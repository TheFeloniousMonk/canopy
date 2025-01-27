import Navigation from "@/components/navbar";
import { AccountWithTxs, Keystore, Validator } from "@/components/api";
import { useEffect, useState } from "react";
import Accounts from "@/components/account";
import Dashboard from "@/components/dashboard";
import Governance from "@/components/governance";
import { Spinner } from "react-bootstrap";

export default function Home() {
  const [state, setState] = useState({ navIdx: 0, keystore: null, keyIdx: 0, account: {}, validator: {} });
  const setNavIdx = (i) => setState({ ...state, navIdx: i });

  function queryAPI(i = state.keyIdx) {
    Keystore().then((ks) => {
      let mergedKS = {...ks.addressMap, ...ks.nicknameMap}

      if (Object.keys(mergedKS).length === 0) {
        console.warn("mergedKS is empty. No data to query.");
        setState({ ...state, keystore: {}, account: {}, validator: {} }); // Handle empty case
        return
      }

      Promise.all([AccountWithTxs(0, mergedKS[Object.keys(mergedKS)[i]].keyAddress, 0), Validator(0, mergedKS[Object.keys(mergedKS)[i]].keyAddress)]).then((r) => {
        setState({ ...state, keyIdx: i, keystore: mergedKS, account: r[0], validator: r[1] });
      });
    });
  }

  useEffect(() => {
    const i = setInterval(() => {
      queryAPI();
    }, 4000);
    return () => clearInterval(i);
  });

  if (state.keystore === null) {
    queryAPI();
    return <Spinner id="spinner" />;
  }
  return (
    <>
      <div id="container">
        <Navigation {...state} setActiveKey={queryAPI} setNavIdx={setNavIdx} />
        <div id="pageContent">
          {(() => {
            if (state.navIdx === 0) {
              return <Accounts keygroup={Object.values(state.keystore)[state.keyIdx]} {...state} />;
            } else if (state.navIdx === 1) {
              return <Governance keygroup={Object.values(state.keystore)[state.keyIdx]} {...state} />;
            } else {
              return <Dashboard />;
            }
          })()}
        </div>
      </div>
    </>
  );
}
