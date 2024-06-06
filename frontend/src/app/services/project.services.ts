import {Injectable} from "@angular/core";
import {HttpClient, HttpParams} from "@angular/common/http";
import {Observable} from "rxjs";
import {IRequest} from "../models/request.model";
import {ConfigurationService} from "./configuration.services";

@Injectable({
    providedIn: 'root'
})
export class ProjectServices {
  urlPath = ""

  constructor(private http: HttpClient, private configurationService: ConfigurationService) {
    this.urlPath = configurationService.getValue("pathUrl")
  }


  getAll(page: number, searchName: string): Observable<IRequest> {
    const params = new HttpParams()
        .set('page', page.toString())
        .set('searchName', searchName);

    return this.http.get<IRequest>(`${this.urlPath}/api/v1/projects`, { params });
}

addProject(key: String): Observable<IRequest> {
    return this.http.post<IRequest>(`${this.urlPath}/api/v1/projects`, { key });
}

deleteProject(id: Number): Observable<IRequest> {
    return this.http.delete<IRequest>(`${this.urlPath}/api/v1/projects/${id}`);
}
}
