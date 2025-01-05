import About from "./components/about";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/home";
import LoginPage from "./components/login";
import Navbar from "./components/navbar";
import PricingPage from "./components/price";
import Request from "./components/request";
import Test from "./components/test";
import ResetPassword from "./components/resetpassword";
import ResetPage from "./components/password-reset";
import MasterHomePage from "./components/masters/masterHomePage";
import UserHomePage from "./components/users/userHomePage";
import AddUsers from "./components/masters/addusers";
import EditProductInfo from "./components/masters/editProductInfo";
import Error from "./components/error";
import InfoPage from "./components/masters/infoPage";
import UserInfo from "./components/users/personalInfo";
import Success from "./components/success";
import AddPassword from "./components/users/addPassword";
import Reset from "./components/reset";
import { useLocation } from "react-router-dom";
import Otp from "./components/otp";
import NotFound from "./components/notFound";
function App() {
  return (
    <>
      <Router>
        <ConditionalNavbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/about" element={<About />} />
          <Route path="/error" element={<Error />} />
          <Route path="/price" element={<PricingPage />} />
          <Route path="/upgradeproduct" element={<Request />} />
          <Route path="/resetpassword" element={<ResetPassword />} />
          <Route path="/resetcreds" element={<ResetPage />} />
          <Route path="/masterhomepage" element={<MasterHomePage />} />
          <Route path="/userhomepage" element={<UserHomePage />} />
          <Route path="/test" element={<Test />} />
          <Route path="/success" element={<Success />} />
          <Route path="/master/adduser" element={<AddUsers />} />
          <Route
            path="/master/editconfig"
            element={<EditProductInfo />}
          ></Route>
          <Route path="/master/info" element={<InfoPage />} />
          <Route path="/user/info" element={<UserInfo />} />
          <Route path="/user/addpassword" element={<AddPassword />} />
          <Route path="/reset" element={<Reset />} />
          <Route path="/otp" element={<Otp />} />
          <Route path="/not" element={<NotFound />} />
        </Routes>
      </Router>
    </>
  );
  function ConditionalNavbar() {
    const location = useLocation();
    const excludeNavbarRoutes = ["/reset", "/login", "/otp", "/error", "/newpassword", "/not", "/resetcreds", "/resetpassword"];
    if (excludeNavbarRoutes.includes(location.pathname)) {
      return null;
    }
    return <Navbar />;
  }
}

export default App;
