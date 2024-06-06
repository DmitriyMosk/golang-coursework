import {Injectable} from "@angular/core";
import {HttpClient, HttpParams} from "@angular/common/http";
import {Observable} from "rxjs";
import {IRequest} from "../models/request.model";
import {IRequestObject} from "../models/requestObj.model";
import {ConfigurationService} from "./configuration.services";

@Injectable({
  providedIn: 'root'
})
export class DatabaseProjectServices {
  urlPath = ""

  constructor(private http: HttpClient, private configurationService: ConfigurationService) {
    this.urlPath = configurationService.getValue("pathUrl")
  }


  getAll(): Observable<IRequest>{
    return this.http.get<IRequest>(`${this.urlPath}/api/v1/all-projects`);
  }

  getProjectStatByID(id: string): Observable<IRequestObject> {
    return this.http.get<IRequestObject>(`${this.urlPath}/api/v1/project/${id}/statistics`);
  }

  getComplitedGraph(taskNumber: string, projectName: Array<string>): Observable<IRequestObject> {
    const params = new HttpParams().set('taskNumber', taskNumber).set('projectName', projectName.join(','));
    return this.http.get<IRequestObject>(`${this.urlPath}/api/v1/completed-graph`, { params });
  }

  getGraph(taskNumber: string, projectName: string): Observable<IRequestObject> {
    return this.http.get<IRequestObject>(`${this.urlPath}/api/v1/graph`, { params: { taskNumber, projectName } });
  }

  makeGraph(taskNumber: string, projectName: string): Observable<IRequestObject> {
    return this.http.post<IRequestObject>(`${this.urlPath}/api/v1/graph`, { taskNumber, projectName });
  }

  deleteGraphs(projectName: string): Observable<IRequestObject> {
    return this.http.delete<IRequestObject>(`${this.urlPath}/api/v1/graph/${projectName}`);
  }

  isAnalyzed(projectName: string): Observable<IRequestObject>{
    return this.http.get<IRequestObject>(`${this.urlPath}/api/v1/is-analyzed/${projectName}`);
  }

  isEmpty(projectName: string): Observable<IRequestObject>{
    return this.http.get<IRequestObject>(`${this.urlPath}/api/v1/is-empty/${projectName}`);
  }
}
