import { React } from "react";
import { useNavigate } from "react-router-dom";
const PricingPage = () => {
  let navigate = useNavigate();
  const handleClick = (type) => {
    navigate("/requestproduct", { state: { planType: type } });
  };
  return (
    <>
      <section className="min-h-screen flex items-center justify-center pt-16">
        <div className="w-full max-w-6xl rounded-xl p-12">
          <h2 className="text-4xl font-bold text-center text-gray-800 mb-12">
            Our Pricing Plans
          </h2>
          <div className=" flex flex-wrap justify-center gap-3">
            {/* Basic Plan */}
            <div className="bg-gray-100 w-full md:w-1/3 lg:w-1/4 bg-white rounded-lg p-8 outline: none">
              <h3 className="text-2xl font-semibold text-center text-gray-800">
                Basic Plan
              </h3>
              <p className="text-center text-gray-600 my-4">$5.00/month</p>
              <ul className="text-gray-600 mb-6">
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> 4 Users
                </li>
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> 10 Passwords
                  Each
                </li>
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> Basic Support
                </li>
              </ul>
              <button
                onClick={() => handleClick("Basic Plan")}
                className="w-full bg-[#AA7DFF] text-white py-2 rounded-lg hover:bg-[#C49DFF] transition-colors duration-300"
              >
                Choose Plan
              </button>
            </div>

            {/* Pro Plan */}
            <div className="w-full md:w-1/3 lg:w-1/4 bg-[#EADCF0] rounded-lg p-8 outline: none">
              <h3 className="text-2xl font-semibold text-center text-gray-800">
                Pro Plan
              </h3>
              <p className="text-center text-gray-600 my-4">$19.99/month</p>
              <ul className="text-gray-600 mb-6">
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> 50 Users
                </li>
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> 50 Passwords
                  Each
                </li>
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> Priority
                  Support
                </li>
              </ul>
              <button
                onClick={() => handleClick("Pro Plan")}
                className="w-full bg-[#AA7DFF] text-white py-2 rounded-lg hover:bg-[#C49DFF] transition-colors duration-300"
              >
                Choose Plan
              </button>
            </div>

            {/* Premium Plan */}
            <div className="bg-gray-100 w-full md:w-1/3 lg:w-1/4  rounded-lg p-8 outline:none">
              <h3 className="text-2xl font-semibold text-center text-gray-800">
                Premium Plan
              </h3>
              <p className="text-center text-gray-600 my-4">$29.99/month</p>
              <ul className="text-gray-600 mb-6">
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> Unlimited
                  Projects
                </li>
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> Unlimited
                  Passwords
                </li>
                <li className="flex items-center mb-2">
                  <span className="mr-2 text-green-500">✔</span> 24/7 Support
                </li>
              </ul>
              <button
                onClick={() => handleClick("Premium Plan")}
                className="w-full bg-[#AA7DFF] text-white py-2 rounded-lg hover:bg-[#C49DFF] transition-colors duration-300"
              >
                Choose Plan
              </button>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default PricingPage;
