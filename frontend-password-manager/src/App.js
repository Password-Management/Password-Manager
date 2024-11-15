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
import UserHomePage from "./components/users/userHomePage"
import AddUsers from "./components/masters/addusers";
import EditProductInfo from "./components/masters/editProductInfo";
import Error from "./components/error";
import InfoPage from "./components/masters/infoPage";
import UserInfo from "./components/users/personalInfo";
import Success from "./components/success";
import AddPassword from "./components/users/addPassword";
function App() {
  return (
    <>
      <Router>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/about" element={<About />} />
          <Route path="/error" element={<Error />} />
          <Route path="/price" element={<PricingPage />} />
          <Route path="/requestproduct" element={<Request />} />
          <Route path="/resetpassword" element={<ResetPassword />} />
          <Route path="/newpassword" element={<ResetPage />} />
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
        </Routes>
      </Router>
    </>
  );
}

export default App;
