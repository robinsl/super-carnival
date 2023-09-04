CREATE MIGRATION m16m2llbuinucivwijsi7tsfmmgmc2iehjjhltavpgbv53godh43za
    ONTO initial
{
  CREATE TYPE default::Department {
      CREATE REQUIRED PROPERTY name: std::str;
  };
  CREATE TYPE default::JobTitle {
      CREATE REQUIRED PROPERTY name: std::str;
  };
  CREATE TYPE default::Employee {
      CREATE LINK job_title: default::JobTitle;
      CREATE REQUIRED PROPERTY birthday: cal::local_date;
      CREATE REQUIRED PROPERTY first_name: std::str;
      CREATE REQUIRED PROPERTY last_name: std::str;
  };
  ALTER TYPE default::Department {
      CREATE MULTI LINK employees: default::Employee;
  };
  ALTER TYPE default::Employee {
      CREATE MULTI LINK departements := (.<employees[IS default::Department]);
  };
};
