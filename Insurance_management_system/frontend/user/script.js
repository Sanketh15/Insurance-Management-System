document.addEventListener("DOMContentLoaded", async function () {
  const username = sessionStorage.getItem("username");

  if (!username) {
    window.location.href = '../login.html'; // Redirect to login if no session
    return;
  }

  history.pushState(null, null, window.location.href); // Prevent going back

  window.onpopstate = function () {
    destroySession();  // Clear session
    window.location.href = '../login.html';  // Redirect to login
  };

  function destroySession() {
    sessionStorage.removeItem('username');
    sessionStorage.removeItem('email');
    sessionStorage.removeItem('userId');
  }

  // Fetch and display user data if username exists
  if (username) {
    document.getElementById("name").value = username;

    try {
      const response = await fetch("http://localhost:8080/users");
      if (response.ok) {
        const users = await response.json();
        const user = users.find(user => user.username === username);
        if (user) {
          document.getElementById("email").value = user.email || "";
        } else {
          console.warn("User not found for username:", username);
        }
      } else {
        console.error("Failed to fetch users. Response:", await response.text());
      }
    } catch (error) {
      console.error("Error occurred while fetching users:", error);
    }
  } else {
    console.error("Username not found in sessionStorage.");
  }

  // Sidebar elements
  const sidebar = document.querySelector('.sidebar');
  const hamburgerBtn = document.querySelector('.hamburger-btn');
 
  // Check and set the sidebar state on page load based on localStorage
  function checkSidebarState() {
    const isShrunk = localStorage.getItem("sidebarShrunk") === "true";
    if (isShrunk) {
      sidebar.classList.add('shrink');
    } else {
      sidebar.classList.remove('shrink');
    }
  }
 

  checkSidebarState();

  hamburgerBtn.addEventListener('click', function () {
    sidebar.classList.toggle('shrink');
    const isShrunk = sidebar.classList.contains('shrink');
    localStorage.setItem("sidebarShrunk", isShrunk.toString());
  });

  const sidebarLinks = document.querySelectorAll('.sidebar .sidebar-content a');
  sidebarLinks.forEach(link => {
    link.addEventListener('click', function () {
      sidebar.classList.add('shrink');
      localStorage.setItem("sidebarShrunk", "true");
    });
  });


  // Logout functionality
  const logoutLink = document.getElementById('logoutlink');
  if (logoutLink) {
    logoutLink.addEventListener('click', function (event) {
      event.preventDefault();
      destroySession(); 
      window.location.href = '../login.html'; 
    });
  }

  // Handle page show and ensure user is logged out if session doesn't exist
  window.onpageshow = function(event) {
    if (!sessionStorage.getItem('username') && !event.persisted) {
      window.location.href = '../login.html'; 
    }
  };
});
