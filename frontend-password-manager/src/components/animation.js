// This is the animation render
import * as assets from "../assets/Login.json";
import * as decryptAssets from "../assets/decrypt.json";
import * as deleteAssets from "../assets/delete.json";
import * as loaderAssets from "../assets/loader.json";
import { Player } from "@lottiefiles/react-lottie-player";

export const LoginAnimation = () => {
  return (
    <>
      <div className="flex flex-col items-center">
        <Player
          autoplay
          loop
          src={assets}
          style={{
            height: "500px",
            width: "500px",
          }}
        />
        <span className="mt-4 text-xl text-gray-700 text-bold">
          Getting things ready for you...
        </span>
      </div>
    </>
  );
};

export const DecryptAnimation = () => {
  return (
    <>
      <div className="flex flex-col items-center">
        <Player
          autoplay
          loop
          src={decryptAssets}
          style={{
            height: "500px",
            width: "500px",
          }}
        />
        <span className="mt-4 text-xl text-black-700 text-bold">
          Decrypting Passwords .....
        </span>
      </div>
    </>
  );
};

export const DeleteAnimation = () => {
  return (
    <>
      <div className="flex flex-col items-center">
        <Player
          autoplay
          loop
          src={deleteAssets}
          style={{
            height: "500px",
            width: "500px",
          }}
        />
        <span className="mt-4 text-xl text-black-700 text-bold">
          Deleting selection Entry
        </span>
      </div>
    </>
  );
};

export const LoaderAnimation = () => {
  return (
    <>
      <div className="flex flex-col items-center">
        <Player
          autoplay
          loop
          src={loaderAssets}
          style={{
            height: "500px",
            width: "500px",
          }}
        />
        <span className="mt-4 text-xl text-black-700 text-bold">
          Deleting selection Entry
        </span>
      </div>
    </>
  );
};
