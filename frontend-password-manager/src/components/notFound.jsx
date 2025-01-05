import { React, useEffect } from "react";
import ImageTest from "../assets/404.jpg";
import { Link, useNavigate } from "react-router-dom";
const NotFound = () => {
  let navigate = useNavigate();
  useEffect(() => {
    const timer = setTimeout(() => {
      navigate("/");
    }, 3000);
    return () => clearTimeout(timer);
  }, []);
  return (
    <>
      <section className="min-h-screen flex items-center justify-center">
        <div className="flex flex-col md:flex-row rounded-2xl max-w-4xl p-5 md:p-10 items-center">
          <div className="w-full md:w-1/2 hidden md:block">
            <img
              className="rounded-2xl object-cover"
              src={ImageTest}
              alt="404-Not-Found"
            />
          </div>
          <div className="w-full md:w-1/2 px-6 md:px-8">
            <h2 className="font-bold text-8xl text-black md:text-left text-center align-items-center">
              404
            </h2>
            <br />
            <h2 className="font-bold text-3xl text-black text-center md:text-left">
              Page Not Found
            </h2>
            <br />
            <h3>
              We're sorry, but the page you're looking for doesn't exist or may
              have been moved.
            </h3>
            <br />
            <Link to={"/"}>
              <button className="bg-[#AA7DFF] text-white px-10 py-2 rounded-full hover:bg-[#C49DFF]">
                Go To Home
              </button>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
};

export default NotFound;
