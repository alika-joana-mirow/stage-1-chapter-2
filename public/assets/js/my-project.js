// getting data
let projects = [];

function getData(event) {
  // prevent from reloading the page
  event.preventDefault();

  let projectName = document.getElementById("project-name").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let desc = document.getElementById("description").value;
  let imageProject = document.getElementById("inputImageProject").files;

  //   image input validation
  imageProject = URL.createObjectURL(imageProject[0]);

  // checking checkbox value
  checkedValue = [];
  let techProject = document.getElementsByClassName("tech");
  let data = techProject.length;
  for (let i = 0; i < data; i++) {
    if (techProject[i].checked == true) {
      checkedValue.push(techProject[i].value);
    }
  }

  let project = {
    projectName,
    desc,
    checkedValue,
    imageProject,
    startDate,
    endDate,
  };

  // add new item
  projects.push(project);
  console.log(projects);

  showData();
}

function showData() {
  document.getElementById("list-show").innerHTML = "";
  for (let i = 0; i < projects.length; i++) {
    document.getElementById("list-show").innerHTML += `
        <div class="list-container">
            <div class="img-list-container">
              <img src="${projects[i].imageProject}" alt="photo" />
            </div>
            <div class="list-content">
              <a href="/article.html">${projects[i].projectName}</a>
              <p class="duration">Duration : ${getDuration(
                projects[i].startDate,
                projects[i].endDate,
              )}</p>
              <p>
                ${projects[i].desc}
              </p>
              <div class="icon-list">
                ${(function icon() {
                  let string = "";
                  for (let j = 0; j < projects[i].checkedValue.length; j++) {
                    string += `<div class="icon-item">
                <i class="${projects[i].checkedValue[j]}"></i>
              </div>`;
                  }

                  return string;
                })()}
              </div>
              <div class="list-action">
                <a href="#" class="edit-action">edit</a>
                <a href="#" class="delete-action">delete</a>
              </div>
            </div>
        </div>`;
  }
}

function getDuration(start, end) {
  let proStart = new Date(start);
  let proEnd = new Date(end);

  let duration = proEnd - proStart;

  let second = 1000;
  let minute = 60;
  let hour = 60;
  let day = 24;
  let week = 7;
  let month = 30;

  let monthDuration = Math.floor(
    duration / (second * minute * hour * day * month),
  );
  if (monthDuration != 0) {
    return monthDuration + " month";
  } else {
    let weekDuration = Math.floor(
      duration / (second * minute * hour * day * week),
    );
    if (weekDuration != 0) {
      return weekDuration + "weeks";
    } else {
      let dayDuration = Math.floor(duration / (second * minute * hour * day));
      if (dayDuration != 0) {
        return dayDuration + "days";
      } else {
        let hoursDuration = Math.floor(duration / (second * minute * hour));
        if (hoursDuration != 0) {
          return hoursDuration + "hours ago";
        } else {
          let MinutesDuration = Math.floor(duration / (second * minute));
          if (MinutesDuration != 0) {
            return MinutesDuration + "minutes ago";
          } else {
            let secondsDuration = Math.floor(duration / second);
            if (secondsDuration != 0) {
              return secondsDuration + "seconds";
            }
          }
        }
      }
    }
  }
}
