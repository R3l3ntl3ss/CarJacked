import React from "react";

import styles from "./dashboard.module.scss";

function Error() {
  return (
    <div className={styles.error}>
      <h1>Error</h1>
      <img src={require("../../images/error.png")} alt="Error" />
      <h2>Something went wrong. Please Try Again</h2>
      <a href="/dashboard">Reload</a>
    </div>
  );
}

export default Error;
