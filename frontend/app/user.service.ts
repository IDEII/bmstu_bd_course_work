import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private _isLoggedIn: boolean = false;
  private admin: boolean = false;
  private apiUrl = 'http://localhost:8080';
  private user_id : string = "";

  constructor(private http: HttpClient) { }

  login_on_site (user_id: string){
    this.user_id = user_id;
    this._isLoggedIn = true;
  }
  logout_on_site(){
    this._isLoggedIn = false;
    this.user_id = "";
  }
  isLoggedIn(): boolean {
    return this._isLoggedIn;
  }
  get_user_id (): string {
    return this.user_id
  }

  register(user: any): Observable<any> {
    const registerHeader = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/registration`, JSON.stringify(user), {headers: registerHeader}).pipe(catchError(this.handleError));
  }
  
  getUserbyE_mail(e_mail: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/user-profile/${e_mail}/email`).pipe(
      map((data: any) => {
        return data.user_id;
      }),
      catchError((error) => {
        console.error('Ошибка при получении пользователя по имени:', error);
        return throwError(() => new Error('Ошибка при получении пользователя'));
      })
    );;
  }

  getUserbyUsername(username: string): Observable<string> {
    return this.http.get<any>(`${this.apiUrl}/user-profile/${username}/username`).pipe(
      map((data: any) => {
        return data.user_id;
      }),
      catchError((error) => {
        console.error('Ошибка при получении пользователя по имени:', error);
        return throwError(() => new Error('Ошибка при получении пользователя'));
      })
    );
  }

  login(data: any): Observable<any> {
    const loginHeader = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/login`, JSON.stringify(data), {headers: loginHeader}).pipe(catchError(this.handleError));
  }

  getUser(id: string): Observable<any>  {
    return this.http.get<any>(`${this.apiUrl}/user-profile/${id}`).pipe(
      map((data: any) => {
        return data.user;
      }),
      catchError((error) => {
        console.error('Ошибка при получении пользователя по имени:', error);
        return throwError(() => new Error('Ошибка при получении пользователя'));
      })
    );;
  }

  updateUser(id:string, user: any): Observable<any>  {
    const updateUserHeader = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/user-profile/${id}/update`, JSON.stringify(user), { headers: updateUserHeader}).pipe(catchError(this.handleError));
  }

  private handleError(error: any): Observable<never> {
    console.error('An error occurred:', error);
    return throwError(() => new Error(error.message || 'Server error'));
  }

  deleteUser(id:string): Observable<any>  {
    return this.http.get<any>(`${this.apiUrl}/user-profile/${id}/delete`);
  }
}
