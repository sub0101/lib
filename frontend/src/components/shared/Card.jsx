import React from 'react';
import style from "../../style/card.module.css"

const Card = () => {
  return (
    <>
      <div className={style.card}>
        <div className={style.count}>1,235</div>
        <div className={style.label}>Total Users</div>
      </div>

      
    </>
  );
}

export default Card;
