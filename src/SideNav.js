import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUser } from "@fortawesome/free-regular-svg-icons";
import { faListAlt } from "@fortawesome/free-regular-svg-icons";
import { faBookmark } from "@fortawesome/free-regular-svg-icons";
import { faExternalLinkAlt } from "@fortawesome/free-solid-svg-icons";
import { faChartBar } from "@fortawesome/free-regular-svg-icons";
import { faQuestionCircle } from "@fortawesome/free-regular-svg-icons";

const SideNav = ({ toggle }) => {
  return (
    <section
      className={`fixed  overflow-scroll top-0 bottom-0 bg-Body h-full text-white  w-11/12  items-center flex-col border-r-[1px] border-DarkGray z-50`}
    >
      {/* Users Info */}
      <article className="w-full px-3 py-2 border-b border-DarkGray ">
        <div className="flex w-full justify-between py-2 px-1">
          <img
            src="https://www.gravatar.com/avatar/1b8fabaa8d66250a7049bdb9ecf44397?s=250&d=mm&r=x"
            alt="#"
            className="w-[55px] h-[55px] rounded-full hover:opacity-90 cursor-pointer lg:hidden sm:flex"
          />
          <button onClick={toggle}>
            <ion-icon name="close-outline" size="large"></ion-icon>
          </button>
        </div>
        <div className="px-1">
          <div className="mx-[5px]">
            <h4 className="font-bold">Richard</h4>
            <p className="text-DarkGray">@richard_77</p>
          </div>
          <div className="flex py-2 px-1">
            <p className="font-bold ">
              8 <span className="font-normal text-DarkGray"> Following</span>
            </p>
            <p className="font-bold ml-2">
              123K <span className="font-normal text-DarkGray"> Followers</span>{" "}
            </p>
          </div>
        </div>
      </article>
      {/* Main Nav  1*/}
      <article className="w-full border-b  border-DarkGray">
        <div className="cursor-pointer flex my-[3px] items-center p-[14px] hover:bg-DarkGray  hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            <FontAwesomeIcon icon={faUser}></FontAwesomeIcon> Profile
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            <FontAwesomeIcon icon={faListAlt}></FontAwesomeIcon> List
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            {" "}
            <ion-icon name="chatbubble-ellipses-outline"></ion-icon> Topics
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            <FontAwesomeIcon icon={faBookmark}></FontAwesomeIcon> Bookmarks
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray  hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            {" "}
            <ion-icon name="flash-outline"></ion-icon> Moments
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150 border-b border-DarkGray ">
          {" "}
          <p>
            {" "}
            <ion-icon name="newspaper-outline"></ion-icon> Newsletter
          </p>
        </div>

        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            {" "}
            <ion-icon name="rocket-outline"></ion-icon> Twitter for
            Professionals
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150">
          {" "}
          <p>
            {" "}
            <FontAwesomeIcon icon={faExternalLinkAlt}></FontAwesomeIcon> Twitter
            Ads
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150 border-b border-DarkGray ">
          {" "}
          <p>
            {" "}
            <FontAwesomeIcon icon={faChartBar}></FontAwesomeIcon> Analytics
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150 ">
          {" "}
          <p>
            {" "}
            <ion-icon name="settings-outline"></ion-icon> Settings
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150  ">
          {" "}
          <p>
            {" "}
            <FontAwesomeIcon icon={faQuestionCircle}></FontAwesomeIcon> Help
            Center
          </p>
        </div>
        <div className="cursor-pointer p-[14px] items-center my-[3px] hover:bg-DarkGray hover:bg-opacity-20 duration-150  ">
          {" "}
          <p> Logout</p>
        </div>
      </article>
    </section>
  );
};

export default SideNav;
