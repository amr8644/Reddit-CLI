import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHome } from "@fortawesome/free-solid-svg-icons";
import { faHashtag } from "@fortawesome/free-solid-svg-icons";
import { faBell } from "@fortawesome/free-regular-svg-icons";
import { faEnvelope } from "@fortawesome/free-regular-svg-icons";
import { faUser } from "@fortawesome/free-regular-svg-icons";
import { faEllipsisH } from "@fortawesome/free-solid-svg-icons";

const Nav = () => {
  return (
    <section className="text-white h-screen  w-[330px] flex items-center justify-between flex-col border-r-[1px] border-DarkGray">
      <article className="w-2/3 h-auto  flex items-start text-2xl  flex-col ">
        <div className="text-white p-[14px]">
          <ion-icon name="logo-twitter" size="large"></ion-icon>
        </div>
        <div className="cursor-pointer flex my-[3px] p-[14px] hover:bg-DarkGray  hover:bg-opacity-20 rounded-full duration-150">
          {" "}
          <p>
            <FontAwesomeIcon icon={faHome}></FontAwesomeIcon> Home
          </p>
        </div>
        <div className="cursor-pointer p-[14px] my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 rounded-full duration-150">
          {" "}
          <p>
            <FontAwesomeIcon icon={faHashtag}></FontAwesomeIcon> Explore
          </p>
        </div>
        <div className="cursor-pointer p-[14px] my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 rounded-full duration-150">
          {" "}
          <p>
            {" "}
            <FontAwesomeIcon icon={faBell}></FontAwesomeIcon> Notifications
          </p>
        </div>
        <div className="cursor-pointer p-[14px] my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 rounded-full duration-150">
          {" "}
          <p>
            <FontAwesomeIcon icon={faEnvelope}></FontAwesomeIcon> Messages{" "}
          </p>
        </div>
        <div className="cursor-pointer p-[14px] my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 rounded-full duration-150">
          {" "}
          <p>
            {" "}
            <FontAwesomeIcon icon={faUser}></FontAwesomeIcon> Profile{" "}
          </p>
        </div>
        <div className="cursor-pointer p-[14px] my-[3px] hover:bg-DarkGray hover:bg-opacity-20 rounded-full duration-150">
          {" "}
          <p>
            {" "}
            <FontAwesomeIcon icon={faEllipsisH}></FontAwesomeIcon> More
          </p>
        </div>
        <button className="px-20 py-4 bg-Primary rounded-full hover:bg-opacity-70 my-[3px]">
          Tweet
        </button>
      </article>
      <div className="cursor-pointer w-4/5 py-3 my-4 hover:bg-DarkGray hover:bg-opacity-20 rounded-full duration-150 flex items-center  justify-evenly">
        <img
          src="https://www.gravatar.com/avatar/1b8fabaa8d66250a7049bdb9ecf44397?s=250&d=mm&r=x"
          alt="#"
          className="mr-[5px] w-[55px] h-[55px] rounded-full"
        />
        <div className="mx-[5px]">
          <h4 className="font-bold">Richard</h4>
          <p className="text-DarkGray">@richard_77</p>
        </div>
        <FontAwesomeIcon icon={faEllipsisH}></FontAwesomeIcon>
      </div>
    </section>
  );
};

export default Nav;
