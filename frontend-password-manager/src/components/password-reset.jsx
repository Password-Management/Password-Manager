import { React, useState } from "react";

const ResetPage = () => {
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [passwordLengthValid, setPasswordLengthValid] = useState(false);

  const toggleConfirmPasswordVisibility = () => {
    setShowConfirmPassword((prevState) => !prevState);
  };

  const handlePasswordChange = (e) => {
    const password = e.target.value;
    setNewPassword(password);
    setPasswordLengthValid(password.length >= 8); // Ensure password is at least 8 characters long
  };

  const handleConfirmPasswordChange = (e) => {
    setConfirmPassword(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!passwordLengthValid) {
      setError("Password must be at least 8 characters long.");
    } else if (newPassword !== confirmPassword) {
      setError("Passwords do not match.");
    } else {
      setError("");
      console.log("Form Submitted");
    }
  };

  return (
    <>
      <section className="flex flex-col h-screen items-center justify-center mt-5">
        <h1 className="text-5xl font-semibold text-center mb-6">
          Set New Password
        </h1>

        <form
          onSubmit={handleSubmit}
          className="flex flex-col gap-4 w-full max-w-sm items-center"
        >
          <input
            className={`p-2 rounded-xl border focus:outline-none focus:ring focus:border-blue-500 w-full ${
              !passwordLengthValid && newPassword ? "border-red-500" : ""
            }`}
            type="email"
            name="email"
            placeholder="New Password"
            value={newPassword}
            onChange={handlePasswordChange}
          />

          <div className="relative w-full">
            <input
              className={`p-2 rounded-xl border w-full focus:outline-none focus:ring focus:border-blue-500 ${
                confirmPassword && confirmPassword !== newPassword
                  ? "border-red-500"
                  : ""
              }`}
              type={showConfirmPassword ? "text" : "password"}
              name="confirmPassword"
              placeholder="Confirm Password"
              value={confirmPassword}
              onChange={handleConfirmPasswordChange}
            />
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="gray"
              className="bi bi-eye absolute top-1/2 right-3 -translate-y-1/2 cursor-pointer"
              viewBox="0 0 16 16"
              onClick={toggleConfirmPasswordVisibility}
            >
              <path d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8zM1.173 8a13.133 13.133 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5c2.12 0 3.879 1.168 5.168 2.457A13.133 13.133 0 0 1 14.828 8c-.058.087-.122.183-.195.288-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5c-2.12 0-3.879-1.168-5.168-2.457A13.134 13.134 0 0 1 1.172 8z" />
              <path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5zM4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0z" />
            </svg>
          </div>

          {error && <p className="text-red-500 text-sm">{error}</p>}

          <button
            type="submit"
            className="bg-[#AA7DFF] rounded-xl text-white py-2 w-full hover:scale-105 transition-transform duration-300"
          >
            Submit
          </button>
        </form>
      </section>
    </>
  );
};

export default ResetPage;
