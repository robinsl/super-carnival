module default {
    type Employee {
        required first_name: str;
        required last_name: str;
        required birthday: cal::local_date;
        job_title: JobTitle;
        multi link departements := .<employees[is Department]
    }

    type JobTitle {
        required name: str;
    }

    type Department {
        required name: str;
        multi employees: Employee;
    }
}
